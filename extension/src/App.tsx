import { useEffect, useState } from "react";
import cls from "./App.module.scss";
import appIcon from './assets/ytb2mpv.svg'
import { ClientGetStatus, WSResponse } from "./services/types";

enum Status {
    Loading = 0,
    Ok = 1,
    Error = 2,
}

export default function App() {

    const [status, setStatus] = useState<Status>(Status.Loading);
    const [daemonStatus, setDaemonStatus] = useState<ClientGetStatus | null>(null);


    useEffect(() => {
        chrome.runtime.sendMessage({ getStatus: true });

        const handler = (message: any & { type: string }) => {
            switch (message.type) {
                case "status": {
                    const s: WSResponse<ClientGetStatus> = message;
                    setDaemonStatus(s);
                    setStatus(message.connected ? Status.Ok : Status.Error);
                    break;
                }
            }
        }
        chrome.runtime.onMessage.addListener(handler);

        return () => {
            chrome.runtime.onMessage.removeListener(handler);
        }
    }, []);
    return (
        <div className={cls.Main}>
            <div className={cls.Status} data-status={status === Status.Ok ? "success" : status === Status.Error ? "error" : "loading"}>
                <img src={appIcon} alt="ytb2mpv" className={cls.Icon} />
                <div className={cls.Text}>
                    {status === Status.Loading ? "Loading..." : status === Status.Ok ? "ytb2mpv is running" : "ytb2mpv server is not running"}
                </div>
            </div>
            <div className={cls.Content}>
                <div className={cls.NotRunning}>
                    <p>Make sure you have ytb2mpv server running. If you didn't, please run it or download and install from&nbsp;
                        <a href="https://github.com/michioxd/ytb2mpv" target="_blank">official respository</a>
                    </p>
                </div>
            </div>
        </div>
    )
}