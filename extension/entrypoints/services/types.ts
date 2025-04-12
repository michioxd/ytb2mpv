export type WSResponse<T extends Record<string, any>> = {
    type: string;
} & T;

export interface ServerResponseInfo {
    mpv_status: number;
    ytdlp_status: number;
    server_version: string;
    mpv_version: string;
    ytdlp_version: string;
}

export interface ClientGetStatus {
    connected: boolean;
    serverVersion: string;
    mpvStatus: number;
    ytdlpStatus: number;
    mpvVersion: string;
    ytdlpVersion: string;
}