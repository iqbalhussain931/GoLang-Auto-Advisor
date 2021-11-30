const cacheName = "app-" + "1.0.0";

self.addEventListener("install", event => {
  console.log("installing app worker 1.0.0");

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
          "/GoLang-Auto-Advisor/web/css/bootstrap.min.css",
          "/GoLang-Auto-Advisor/web/css/style.css",
          "/GoLang-Auto-Advisor/web/js/bootstrap.min.js",
          "/GoLang-Auto-Advisor/web/js/script.js",
          "https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css",
          "https://fonts.googleapis.com/css?family=Raleway",
          "https://kit.fontawesome.com/8ed9c3141e.js",
          "https://storage.googleapis.com/murlok-github/icon-192.png",
          "https://storage.googleapis.com/murlok-github/icon-512.png",
          "https://www.w3schools.com/w3css/4/w3.css",
          
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
  console.log("app worker 1.0.0 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
