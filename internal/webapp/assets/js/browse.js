const componentBrowse = {
    template: `
      <div class="row">
        <h3>Browse</h3>
        <div class="row">
        <div class="col-3">
            <ul class="list-group">
                <span v-for="(files,sessionName) in groupedFiles">
                    <li class="list-group-item list-group-item-dark" aria-current="true">
                        <h5>{{sessionName}}</h5>
                    </li>
                    <li v-for="file in files" class="list-group-item"  v-bind:class="{ active: isItemActive(sessionName,file) }" @click="selectItem(sessionName,file)">{{file}}</li>
                </span>
            </ul>   
        </div>
        <div class="col-9">
        <pre>
            {{fileContent}}
        </pre>
        </div>
        </div>
      </div>`,
    data() {
        return {
            groupedFiles: [],
            selectedItem: {
                session: "",
                file: "",
            },
            fileContent: "",
        }
    },
    methods: {
        async selectItem(sessionName, file) {
            this.selectedItem.session = sessionName;
            this.selectedItem.file = file;
            await this.loadFileContent();
        },
        isItemActive(sessionName, file) {
            return this.selectedItem.session === sessionName && this.selectedItem.file === file;
        },
        async loadFileContent() {
            let data = await getFile(this.selectedItem.session, this.selectedItem.file)
            this.fileContent = this.b64_to_utf8(data.content);
        },
        b64_to_utf8(str) {
            return decodeURIComponent(escape(window.atob(str)));
        }
    },
    async mounted() {
        let groupedFiles = {}
        let data = await getFiles("")

        data.files.forEach((e) => {
            if (!groupedFiles[e.session]) groupedFiles[e.session] = [];

            groupedFiles[e.session].push(e.file)
        })

        this.groupedFiles = groupedFiles;
    }
}