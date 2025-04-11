export function extractYouTubeId(url: string): string | null {
    const regex = /https?:\/\/(?:www\.)?youtube\.com\/watch\?v=([0-9A-Za-z_-]{1,10})$/;
    const match = url.match(regex);
    return match ? match[1] : null;
}
