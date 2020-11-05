import EditorJS from '@editorjs/editorjs'; 
import Header from '@editorjs/header'; 
import List from '@editorjs/list'; 
import ImageTool from '@editorjs/image';

const backendUrl = "http://localhost:8080/medias"
const backendAssetUrl = "http://localhost:8080/assets/media/"
// const uploadHeaders = {
//   "Content-Type": "multipart/form-data",
// }

const editor = new EditorJS({ 
  holder: 'editorjs', 
  placeholder: 'Let`s write an awesome story!',

  tools: {
    header: Header,
    list: List,
    image: {
      class: ImageTool,
      config: {
        uploader: {
          uploadByFile(file) {
            const formData  = new FormData();
            formData.append("files", file);
            console.log(file)
            return fetch(backendUrl, {method: 'POST', body: formData}).then((res) => {
              console.log(res);
              return {
                file: { url: backendAssetUrl + file.name }
              };
            });
          }
        },
      }
    }
  },
})


// editor.save().then((outputData) => {
//     console.log('Article data: ', outputData)
//   }).catch((error) => {
//     console.log('Saving failed: ', error)
//   });