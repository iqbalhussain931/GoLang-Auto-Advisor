const cacheName = "app-" + "e2e41cf80c2d9f4d1c88848ab20faf8ed3ece705";

self.addEventListener("install", event => {
  console.log("installing app worker e2e41cf80c2d9f4d1c88848ab20faf8ed3ece705");

  event.waitUntil(
    caches.open(cacheName).
      then(cache => {
        return cache.addAll([
          "/GoLang-Auto-Advisor",
          "/GoLang-Auto-Advisor/app.css",
          "/GoLang-Auto-Advisor/app.js",
          "/GoLang-Auto-Advisor/manifest.webmanifest",
          "/GoLang-Auto-Advisor/wasm_exec.js",
          "/GoLang-Auto-Advisor/web/app.wasm",
          "https://storage.googleapis.com/murlok-github/icon-192.png",
          "https://storage.googleapis.com/murlok-github/icon-512.png",
          
        ]);
      }).
      then(() => {
        self.skipWaiting();
      })
  );
});

self.addEventListener("activate", event => {
  event.waitUntil(
    caches.keys().then(keyList => {
      return Promise.all(
        keyList.map(key => {
          if (key !== cacheName) {
            return caches.delete(key);
          }
        })
      );
    })
  );
  console.log("app worker e2e41cf80c2d9f4d1c88848ab20faf8ed3ece705 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
