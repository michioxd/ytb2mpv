export default defineContentScript({
  matches: ['*://*.youtube.com/*'],
  main() {
    console.log('is that working?');
  },
});
