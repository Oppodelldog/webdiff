const componentBrowse = {
    template: `
<div class="row">
    <h3>Browse</h3>
    <div class="row">
        <div class="col-3">
            <ul class="list-group">
                <li class="list-group-item list-group-item-dark" aria-current="true">
                    Options
                </li>
                <li class="list-group-item">
                    <div class="form-check">
                        <input class="form-check-input" type="checkbox" v-model="applySelectedFilter" @change="loadFileContent()" id="chkAutoApplyFilter">
                        <label class="form-check-label" for="chkAutoApplyFilter">
                            apply filter ({{selectedFilterName}}) to selected file
                        </label>
                    </div>
                <div class="form-check">
                        <input class="form-check-input" type="checkbox" v-model="prettifyHtml" @change="loadFileContent()" id="chPrettifyHtml">
                        <label class="form-check-label" for="chPrettifyHtml">
                            prettify html output
                        </label>
                    </div>                    
                </li>
            </ul>
            <ul class="list-group">
                <span v-for="(files,sessionName) in groupedFiles">
                    <li class="list-group-item list-group-item-dark" aria-current="true">
                        <h5>{{sessionName}}</h5>
                    </li>
                    <li v-for="file in files" class="list-group-item"
                        v-bind:class="{ active: isItemActive(sessionName,file) }" @click="selectItem(sessionName,file)">{{file}}</li>
                </span>
            </ul>
        </div>
        <div class="col-9">
            <div class="row">
                <filters 
                    v-model="selectedFilterName"
                    @change="loadFileContent()" 
                    @saved="loadFileContent()"
                    @deleted="loadFileContent()"
                />
                <div class="row mt-3">
                    <strong>Content</strong>
                    <pre>
        
{{fileContent}}

                    </pre>
                </div>
            </div>
        </div>
    </div>
</div>`,
    data() {
        return {
            groupedFiles: [],
            applySelectedFilter: false,
            prettifyHtml: false,
            selectedFilterName: "",
            selectedItem: {
                session: "",
                file: "",
            },
            fileContent: "<-- select a file to load",
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
            const filterName = (this.applySelectedFilter) ? this.selectedFilterName : "";

            if (this.selectedItem.session !== "" && this.selectedItem.file !== "") {
                let data = await getFile(this.selectedItem.session, this.selectedItem.file, filterName, this.prettifyHtml)
                if (data.content === null || data.error !== undefined) {
                    this.fileContent = ""
                    if (data.error !== undefined) {
                        this.fileContent = data.error;
                    }
                } else {
                    this.fileContent = this.b64_to_utf8(data.content);
                }
            }
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