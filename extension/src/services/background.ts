import { ClientGetStatus, ServerResponseInfo, WSResponse } from "./types";

class Ytb2MpvClient {
    public socket: WebSocket | null = null;
    public connected: boolean = false;
    public serverVersion: string = "";
    public mpvStatus: number = 0;
    public ytdlpStatus: number = 0;
    public mpvVersion: string = "";
    public ytdlpVersion: string = "";

    private onOpen: (event: Event) => void = () => {
        this.connected = true;
        this.sendToClient("status");
        console.log("Connected to the server");
    }

    private onMessage: (event: MessageEvent) => void = (event) => {
        const data = JSON.parse(event.data) as WSResponse<any>;

        switch (data.type) {
            case "server_info":
                {
                    const d = data as WSResponse<ServerResponseInfo>;
                    this.serverVersion = d.server_version;
                    this.mpvStatus = d.mpv_status;
                    this.ytdlpStatus = d.ytdlp_status;
                    this.mpvVersion = d.mpv_version;
                    this.ytdlpVersion = d.ytdlp_version;

                    break;
                }
        }
        console.log("Received data: ", data);
    }

    private onError: (event: Event) => void = (error) => {
        console.error("WebSocket error: ", error);
    }

    private onClose: (event: CloseEvent) => void = () => {
        console.log("Disconnected, reconnecting...");
        this.reset();
        setTimeout(() => {
            this.connect();
        }, 1000);
    }

    private reset() {
        this.connected = false;
        this.serverVersion = "";
        this.mpvStatus = 0;
        this.ytdlpStatus = 0;
        this.mpvVersion = "";
        this.ytdlpVersion = "";
        this.sendToClient("status");
    }

    constructor() {
        this.reset();
    }

    public connect() {
        if (this.socket) {
            this.socket.close();
            this.socket?.removeEventListener("open", this.onOpen);
            this.socket?.removeEventListener("message", this.onMessage);
            this.socket?.removeEventListener("error", this.onError);
            this.socket?.removeEventListener("close", this.onClose);
        }
        this.connected = false;
        this.socket = new WebSocket("ws://localhost:53918/ytb2mpv");
        this.socket.addEventListener("open", this.onOpen);
        this.socket.addEventListener("message", this.onMessage);
        this.socket.addEventListener("error", this.onError);
        this.socket.addEventListener("close", this.onClose);
    }

    public sendToClient(type: "status") {
        let data: any = {};
        switch (type) {
            case "status":
                data = {
                    type: "status",
                    connected: this.connected,
                    serverVersion: this.serverVersion,
                    mpvStatus: this.mpvStatus,
                    ytdlpStatus: this.ytdlpStatus,
                    mpvVersion: this.mpvVersion,
                    ytdlpVersion: this.ytdlpVersion
                } as ClientGetStatus;
                break;
        }
        chrome.runtime.sendMessage(data);

    }
}

const ytb2mpvClient = new Ytb2MpvClient();

ytb2mpvClient.connect();

chrome.runtime.onMessage.addListener((message) => {
    if (message.getStatus) {
        ytb2mpvClient.sendToClient("status");
        return;
    }
});