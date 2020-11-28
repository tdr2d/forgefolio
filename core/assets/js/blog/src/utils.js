export function getInitialData(){
  return {
    "time": new Date().getTime(),
    "blocks": [
      {
        "type": "title",
        "data": {
          "text": "My title",
        }
      },
      {
        "type": "paragraph",
        "data": {
          "text": "Let`s write an awesome story!"
        }
      },
      {
        "type": "image",
        "data": {
            "url": "https://images.unsplash.com/photo-1601758174114-e711c0cbaa69?w=500",
            "caption": "My caption"
        }
      }
      // {
      //   "type": "html",
      //   "data": {
      //       "html": "PGJsb2NrcXVvdGUgY2xhc3M9InR3aXR0ZXItdHdlZXQiPjxwIGxhbmc9ImVuIiBkaXI9Imx0ciI+U3Vuc2V0cyBkb24mIzM5O3QgZ2V0IG11Y2ggYmV0dGVyIHRoYW4gdGhpcyBvbmUgb3ZlciA8YSBocmVmPSJodHRwczovL3R3aXR0ZXIuY29tL0dyYW5kVGV0b25OUFM/cmVmX3NyYz10d3NyYyU1RXRmdyI+QEdyYW5kVGV0b25OUFM8L2E+LiA8YSBocmVmPSJodHRwczovL3R3aXR0ZXIuY29tL2hhc2h0YWcvbmF0dXJlP3NyYz1oYXNoJmFtcDtyZWZfc3JjPXR3c3JjJTVFdGZ3Ij4jbmF0dXJlPC9hPiA8YSBocmVmPSJodHRwczovL3R3aXR0ZXIuY29tL2hhc2h0YWcvc3Vuc2V0P3NyYz1oYXNoJmFtcDtyZWZfc3JjPXR3c3JjJTVFdGZ3Ij4jc3Vuc2V0PC9hPiA8YSBocmVmPSJodHRwOi8vdC5jby9ZdUt5MnJjanlVIj5waWMudHdpdHRlci5jb20vWXVLeTJyY2p5VTwvYT48L3A+Jm1kYXNoOyBVUyBEZXBhcnRtZW50IG9mIHRoZSBJbnRlcmlvciAoQEludGVyaW9yKSA8YSBocmVmPSJodHRwczovL3R3aXR0ZXIuY29tL0ludGVyaW9yL3N0YXR1cy80NjM0NDA0MjQxNDE0NTk0NTY/cmVmX3NyYz10d3NyYyU1RXRmdyI+TWF5IDUsIDIwMTQ8L2E+PC9ibG9ja3F1b3RlPiA8c2NyaXB0IGFzeW5jIHNyYz0iaHR0cHM6Ly9wbGF0Zm9ybS50d2l0dGVyLmNvbS93aWRnZXRzLmpzIiBjaGFyc2V0PSJ1dGYtOCI+PC9zY3JpcHQ+Cg==",
      //       "view": "code"
      // }
    ],
    "version": "2.19.0"
  };
}

export function setEditorClasses(config, element) {
  if (config && config.customCssClasses && config.customCssClasses && config.customCssClasses.length > 0) {
    config.customCssClasses.forEach(e => {
      element.classList.add(e);
    });
  }
}


const backendUrl = "medias";
const backendAssetUrl = "assets/media/";
function uploadByFile(file) {
  const formData  = new FormData();
  formData.append("medias", file);
  return fetch(backendUrl, {method: 'POST', body: formData}).then((res) => {
    return {
      success: 1,
      file: { url: location.origin + '/' + backendAssetUrl + file.name }
    };
  });
}