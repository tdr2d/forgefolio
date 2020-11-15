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
      }
    ],
    "version": "2.19.0"
  };
}

// FIXME import from jet
export const EDITORJS_IDENTIFIERS = {
  paragraph: "my-paragraph"
}