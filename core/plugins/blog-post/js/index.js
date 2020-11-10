import EditorJS from '@editorjs/editorjs'; 
import Header from '@editorjs/header'; 
import List from '@editorjs/list'; 
import ImageTool from '@editorjs/image';

const backendUrl = "medias"
const backendAssetUrl = "assets/media/"
// const uploadHeaders = {
//   "Content-Type": "multipart/form-data",
// }

const editor = new EditorJS({ 
  holder: 'editorjs', 
  placeholder: 'Let`s write an awesome story!',

  tools: {
    header: Header,
    image: {
      class: ImageTool,
      config: {
        uploader: {
          uploadByFile(file) {
            const formData  = new FormData();
            formData.append("medias", file);
            return fetch(backendUrl, {method: 'POST', body: formData}).then((res) => {
              return {
                success: 1,
                file: { url: location.origin + '/' + backendAssetUrl + file.name }
              };
            });
          }
        },
      }
    }
  },
})

const saveButton = document.getElementById('save-button');
const output = document.getElementById('output');
saveButton.addEventListener('click', () => {
  editor.save().then( savedData => {
    output.innerHTML = JSON.stringify(savedData, null, 4);
  })
})


// .image-toool__caption
// border: none;
// text-align: center;
// font-size: small;
// editor.save().then((outputData) => {
//     console.log('Article data: ', outputData)
//   }).catch((error) => {
//     console.log('Saving failed: ', error)
//   });