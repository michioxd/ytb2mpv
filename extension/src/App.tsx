import { useEffect, useMemo, useRef, useState } from "react";
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
    const [confirmShutdown, setConfirmShutdown] = useState(false);
    const confirmShutdownTimeout = useRef<number | null>(null);

    useEffect(() => {
        if (confirmShutdown) {
            confirmShutdownTimeout.current = window.setTimeout(() => {
                setConfirmShutdown(false);
            }, 4000);
        } else if (confirmShutdownTimeout.current) {
            window.clearTimeout(confirmShutdownTimeout.current);
            confirmShutdownTimeout.current = null;
        }
    }, [confirmShutdown]);

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

    const statusFull = useMemo(() => {
        if (status === Status.Error)
            return {
                text: "ytb2mpv server is not running",
                status: "error"
            }
        if (daemonStatus?.mpvStatus && daemonStatus.mpvStatus > 0 || daemonStatus?.ytdlpStatus && daemonStatus.ytdlpStatus > 0)
            if (daemonStatus?.mpvStatus && daemonStatus.mpvStatus > 0)
                return {
                    text: "mpv having issues!",
                    status: "warning"
                }
            else
                return {
                    text: "yt-dlp having issues!",
                    status: "warning"
                }
        if (status === Status.Ok)
            return {
                text: "ytb2mpv is running",
                status: "success"
            }
        return {
            text: "Loading...",
            status: "loading"
        }
    }, [status, daemonStatus]);
    return (
        <div className={cls.Main}>
            <div className={cls.Status} data-status={statusFull.status}>
                <img src={appIcon} alt="ytb2mpv" className={cls.Icon} />
                <div className={cls.Text}>
                    {statusFull.text}
                </div>
            </div>
            <div className={cls.Content}>
                {status === Status.Ok ? <div className={cls.Running}>
                    <b>No valid YouTube playing tab is open</b>
                    <p>Please switch to a YouTube tab and play a video.</p>
                    <div className="divider"></div>
                    <div className={cls.Info}>
                        <p>ytb2mpv daemon v{daemonStatus?.serverVersion}</p>
                        <p>mpv: {daemonStatus?.mpvVersion || "Not found"}</p>
                        <p>yt-dlp: {daemonStatus?.ytdlpVersion || "Not found"}</p>
                    </div>
                    <div className={cls.Action}>
                        <button className="btn" onClick={() => {
                            chrome.runtime.sendMessage({ reconnect: true });
                        }}>
                            Reconnect to daemon
                        </button>
                        <button className="btn" onClick={() => {
                            if (!confirmShutdown) {
                                setConfirmShutdown(true);
                                return;
                            }
                            chrome.runtime.sendMessage({ shutdown: true });
                            setConfirmShutdown(false);
                        }}>
                            {confirmShutdown ? "Are you sure? Press again" : "Shut down daemon"}
                        </button>
                    </div>
                </div> :
                    <div className={cls.NotRunning}>
                        <p>Make sure you have ytb2mpv server running. If you didn't, please run it or download and install from&nbsp;
                            <a href="https://github.com/michioxd/ytb2mpv" target="_blank">official respository</a>
                        </p>
                    </div>
                }
            </div>
            <div className={cls.Me}>
                w&#47; &hearts; by michioxd - fork me on&nbsp;<a target="_blank" href="https://github.com/michioxd/ytb2mpv">GitHub</a>
            </div>
        </div>
    )
}