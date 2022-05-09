const componentDiff = {
    template: `
<div class="row">
    <div class="col-sm-2">
        <select v-model="selectedSessionA" @change="loadFilesA()" class="form-select"
                aria-label="Choose a session to load file for selection on the left side">
            <option disabled value="">Choose a session</option>
            <option v-for="session in sessions" :value="session">{{session}}</option>
        </select>
        <select class="form-select" multiple
                aria-label="select files to diff with files selected on the right side" v-model="selectedFilesA"
                @change="tryDiff()">
            <option v-for="file in filesA" :value="file.file">{{file.file}}</option>
        </select>
    </div>
    <div class="col-sm-8">
        <div class="accordion-item">
            <h2 class="accordion-header" id="headingOne">
                <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse"
                        data-bs-target="#collapseOne" aria-expanded="false" aria-controls="collapseOne">
                    Filters
                </button>
            </h2>
            <div id="collapseOne" class="accordion-collapse collapse" aria-labelledby="headingOne"
                 data-bs-parent="#accordionExample">
                <div class="accordion-body">
                    <div class="input-group">
                        <span class="input-group-text">Filter</span>
                        <select class="form-select"  @change="tryDiff()"  aria-label="Default filter" v-model="selectedFilter">
                            <option value="">No Filter</option>
                            <option v-for="filter in filters" :value="filter.name">{{filter.name}}</option>
                        </select>
                    </div>
                    <div class="input-group">
                        <div class="form-check">
                            <input class="form-check-input" type="checkbox" v-model="prettifyHtml"
                                   @change="tryDiff()" id="chPrettifyHtml">
                            <label class="form-check-label" for="chPrettifyHtml">
                                prettify html output
                            </label>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div v-for="d in diffContents">
                    <span class="badge bg-dark">
                        {{d.session_a}}/{{d.id_a}}
                    </span>
            <i class="bi bi-file-diff-fill"></i>
            <span class="badge bg-dark">
                        {{d.session_b}}/{{d.id_b}}
                    </span>
            <pre>{{d.diff}}</pre>
        </div>
    </div>
    <div class="col-sm-2">
        <select v-model="selectedSessionB" @change="loadFilesB()" class="form-select"
                aria-label="Choose a session to load file for selection on the right side">
            <option disabled value="">Choose a session</option>
            <option v-for="session in sessions" :value="session">{{session}}</option>
        </select>
        <select class="form-select" multiple
                aria-label="select files to diff with files selected on the left side" v-model="selectedFilesB"
                @change="tryDiff()">
            <option v-for="file in filesB" :value="file.file">{{file.file}}</option>
        </select>
    </div>
</div>`,
    data() {
        return {
            sessions: [],
            filesA: [],
            filesB: [],
            selectedSessionA: "",
            selectedSessionB: "",
            selectedFilesA: [],
            selectedFilesB: [],
            diffContents: [],
            filters: [],
            selectedFilter: "",
            prettifyHtml: false,
        }
    },
    methods: {
        setSessions(data) {
            this.sessions = data.sessions;
        },
        async loadFilesA() {
            if (this.selectedSessionA.length === 0) return;

            const data = await getFiles(this.selectedSessionA)
            this.filesA = data.files;
            this.selectedFilesA = [];
        },
        async loadFilesB() {
            if (this.selectedSessionB.length === 0) return;

            const data = await getFiles(this.selectedSessionB)
            this.filesB = data.files;
            this.selectedFilesB = [];
        },
        async tryDiff() {
            if (this.selectedSessionA.length === 0) return;
            if (this.selectedSessionB.length === 0) return;
            if (this.selectedFilesA.length === 0) return;
            if (this.selectedFilesB.length === 0) return;
            if (this.selectedFilesA.length !== this.selectedFilesB.length) return;

            this.diffContents = [];
            for (let i = 0; i < this.selectedFilesA.length; i++) {
                const sessionA = this.selectedSessionA;
                const sessionB = this.selectedSessionB;
                const fileA = this.selectedFilesA[i];
                const fileB = this.selectedFilesB[i];
                const filter = this.selectedFilter;
                const prettify = this.prettifyHtml;

                const diffResult = await getDiff(sessionA, fileA, sessionB, fileB, filter, prettify)
                this.diffContents.push(diffResult)
            }
        }
    },
    async mounted() {
        this.setSessions(await getSessions())
        this.filters = (await getFilters()).filters;
    }
}