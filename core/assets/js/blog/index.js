import EditorJS from '@editorjs/editorjs'; 
import Header from '@editorjs/header'; 
import ImageTool from '@editorjs/image';
import {getInitialData} from './initialData';

const backendUrl = "medias";
const backendAssetUrl = "assets/media/";
const baseUrl = "/admin";
const editor = new EditorJS({ 
  holder: 'editorjs',
  data: IS_NEW ? getInitialData() : POST,
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

function findTitle(editorJsData) {
  for (i in editorJsData.blocks) {
    const block = editorJsData.blocks[i];
    if (block.type == "header" || block.type == "paragraph") {
      if (block.data.text && block.data.text.length > 3) {
        return block.data.text.substr(0, 63).replace(' ', '_'); 
      }
    }
  }
  return "";
}

function postBlogPost(body) {
  const method = IS_NEW ? 'POST' : 'PATCH';
  const url = IS_NEW ? `${baseUrl}/blog-posts` : `${baseUrl}/blog-posts/${POST.id}`;
  return fetch(url, {method: method, headers: {'Content-Type': 'application/json'}, body: JSON.stringify(body)})
}

document.getElementById('save-button').addEventListener('click', () => {
  editor.save()
  .then(editorJsData => {
    title = findTitle(editorJsData);
    if (!title) {
      return Promise.reject("Wrong title");
    }
    body = {
      id: `BP__${Number.MAX_SAFE_INTEGER - editorJsData.time}__${title}`,
      time: editorJsData.time,
      version: editorJsData.version,
      blocks: JSON.stringify(editorJsData.blocks)
    };
    document.getElementById('output').innerHTML = JSON.stringify(editorJsData, null, 4);
    return postBlogPost(body).then(() => Promise.resolve(body))
  })
  .then((body, err) => {
    if (err) {
      console.error(err)
    } else {
      if (!IS_NEW) {
        window.location.href = `${baseUrl}/blog-posts/${body.id}?notification=${encodeURIComponent("Post updated")}`;
      } else {
        window.location.href = `${baseUrl}/blog-posts?notification=${encodeURIComponent("New post created")}`;
      }
    }
  })
})