// service-worker.js
// TODO 接入缓存

self.addEventListener('install', (event) => {
  console.log('Service Worker installed');
  event.waitUntil(self.skipWaiting());
});

self.addEventListener('activate', (event) => {
  console.log('Service Worker activated');
  event.waitUntil(self.clients.claim());
});

self.addEventListener('fetch', (event) => {
  console.log('Fetch event:', event.request.url);
  const requestUrl = new URL(event.request.url);

  // 过滤掉 chrome-extensions:// 的请求
  if (requestUrl.protocol === 'chrome-extension:') {
    console.log('Fetch filtered:', event.request.url);
    return; // 不处理该请求
  }

  // 继续处理其他请求
  event.respondWith(
    fetch(event.request)
      .then(response => response)
      .catch(error => {
        console.error('Fetch failed:', error);
      })
  );
});