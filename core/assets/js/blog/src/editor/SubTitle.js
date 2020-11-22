
/* SubTitle renders text in h2 tag */
export default class SubTitle {
    static get toolbox() {
      return {
        title: 'SubTitle',
        icon: '<svg xmlns="http://www.w3.org/2000/svg" enable-background="new 0 0 24 24" height="24" viewBox="0 0 24 24" width="24"><g><rect fill="none" height="24" width="24"/></g><g><g><g><path d="M2.5,4v3h5v12h3V7h5V4H2.5z M21.5,9h-9v3h3v7h3v-7h3V9z"/></g></g></g></svg>'
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
      let tag = document.createElement('h2');
      if (this.config.customCssClasses && this.config.customCssClasses.length > 0) {
        this.config.customCssClasses.forEach(e => {
          tag.classList.add(e);
        });
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