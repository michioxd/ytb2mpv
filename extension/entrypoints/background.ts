import Ytb2MpvClient from "./services/client";

export default defineBackground(() => {
    const ytb2mpvClient = new Ytb2MpvClient();

    ytb2mpvClient.connect();

    browser.runtime.onMessage.addListener((message) => {
        if (message.getStatus) {
            ytb2mpvClient.sendToClient("status");
            return;
        }

        if (message.shutdown) {
            ytb2mpvClient.sendToServer({ type: "shutdown" });
            return;
        }

        if (message.reconnect) {
            ytb2mpvClient.reconnect();
            return;
        }
    });
});

