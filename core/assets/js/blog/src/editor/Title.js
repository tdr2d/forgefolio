
/* Title renders text in h1 tag */
export default class Title {
    static get toolbox() {
      return {
        title: 'Title',
        icon: '<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path d="M5 4v3h5.5v12h3V7H19V4z"/></svg>'
      };
    }

    constructor({data, config, api, readOnly}) {
        this.data = data;
        this.config = config;
        this.api = api;
        this.readOnly = readOnly;
        this._text = data && data.text ? data.text : "";
    }
  
    render(){
      let tag = document.createElement('h1');
      if (this.config.customCss && 'title' in this.config.customCss) {
        tag.classList.add(this.config.customCss.title);
      }
      tag.innerHTML = this._text;
      tag.contentEditable = !this.readOnly;
      return tag;
    }
    
    save(blockContent){
      return {
        text: blockContent.innerText
      }
    }

    validate(data) {
      return data.text.trim() !== '';
    }
  }