import EditorJS from '@editorjs/editorjs'; 
import { getInitialData } from './utils';
import Title from './editor/Title';
import SubTitle from './editor/SubTitle';
import Paragraph from './editor/Paragraph';
import Html from './editor/Html';
import Image from './editor/Image';
import Quote from './editor/Quote';

const baseUrl = BASE_URL;
const classes = THEME_CONFIG.blog.editorClasses;

const editor = new EditorJS({ 
  holder: 'editorjs',
  data: IS_NEW ? getInitialData() : POST,
  logLevel: 'VERBOSE',
  tools: {
    title:    { class: Title,    config: { customCssClasses: classes.h1 ? classes.h1 : [] }},
    subtitle: { class: SubTitle, config: { customCssClasses: classes.h2 ? classes.h2 : [] }},
    paragraph: {
      class: Paragraph,
      inlineToolbar: true,
      config: { customCssClasses: classes.p ? classes.p : [] }
    },
    html: {class: Html, config: { customCssClasses: classes.html ? classes.html : [] }},
    image: { class: Image, inlineToolbar: true, config: {
      css: {
          image: classes.image ? classes.image : [],
          imageHolder: classes.imageHolder ? classes.imageHolder : [],
          imageCaption: classes.imageCaption ? classes.imageCaption : [],
        }
      }
    },
    quote: { class: Quote, inlineToolbar: true, config: {
      css: {
        quote: classes.quote ? classes.quote : [],
        p: classes.p ? classes.p : []
      }
    }}
  }
});

function findTitle(editorJsData) {
  for (i in editorJsData.blocks) {
    const block = editorJsData.blocks[i];
    if (block.type == "title" || block.type == "paragraph") {
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
    // return postBlogPost(body).then(() => Promise.resolve(body))
  })
  // .then((body, err) => {
  //   if (err) {
  //     console.error(err)
  //   } else {
  //     if (!IS_NEW) {
  //       window.location.href = `${baseUrl}/blog-posts/${body.id}?notification=${encodeURIComponent("Post updated")}`;
  //     } else {
  //       window.location.href = `${baseUrl}/blog-posts?notification=${encodeURIComponent("New post created")}`;
  //     }
  //   }
  // })
})