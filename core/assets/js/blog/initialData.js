export function getInitialData(){
  return {
    "time": new Date().getTime(),
    "blocks": [
      {
        "type": "header",
        "data": {
          "text": "My title",
          "level": 1
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