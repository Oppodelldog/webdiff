const componentDownload = {
    template: `
 <div class="row">
    <div class="input-group">
        <div>
            <label for="floatingTextarea2">Session</label>
            <input type="text" class="form-control" placeholder="New Session" aria-label="New Session"
                   v-model="selectedSession">
        </div>
        <div class="ms-3">
            <label for="floatingTextarea2">existing sessions</label>
            <select v-model="selectedSession" class="form-select"
                    aria-label="Choose a session to load file for selection on the left side">
                <option disabled value="">Choose a session</option>
                <option v-for="session in sessions" :value="session">{{session}}</option>
            </select>
        </div>
        <div class="ms-3">
            <label for="selFileIdStrategy">file id strategy</label>
            <select id="selFileIdStrategy" v-model="fileIdStrategy" class="form-select"
                    aria-label="Choose a strategy for creating file ids">
                <option value="pipe">Pipe separated</option>                    
                <option value="gen:hash_url">Hash from url</option>
                <option value="gen:hash_path">Hash from path</option>
            </select>
        </div>
    </div>
</div>
<div class="row mt-3">
    <div class="col">
        <label for="floatingTextarea2">Urls</label>
        <textarea v-model="urls" class="form-control" placeholder="one url per line" id="urlList"
                  style="height: 200px"></textarea>
    </div>
</div>
<div class="row mt-3">
    <div class="col">
        <button type="button" :disabled="isFillUrlsButtonDisabled()" @click="fillUrls()"
                class="btn btn-outline-primary">Fill Urls from
            existing session
        </button>
    </div>
    <div class="col">
    <span v-if="hasErrors()" class="badge bg-danger">{{errorMessage}}</span>
    </div>
    <div class="col-auto">
        <button type="button" class="btn btn-primary" @click="download()">Download</button>
    </div>
</div>
<div class="row mt-3">
    <div class="col">
        <div class="alert alert-success" role="alert" v-if="downloadSuccess">
            Downloads have been enqueued
        </div>
        <div class="alert alert-danger" role="alert" v-if="downloadFailure">
            Error while enqueuing
        </div>
    </div>
</div>

        `,
    data() {
        return {
            sessions: [],
            urls: "",
            selectedSession: "",
            downloadSuccess: false,
            downloadFailure: false,
            fileIdStrategy: "pipe",
            errorMessage: "",
            idUrlSeparator: "|",
        }
    },
    methods: {
        setSessions(data) {
            this.sessions = data.sessions;
        },
        async fillUrls() {
            const data = await getSessionUrls(this.selectedSession)
            this.urls = data.urls.join("\n")
        },
        isFillUrlsButtonDisabled() {
            for (let i = 0; i < this.sessions.length; i++) {
                if (this.sessions[i] === this.selectedSession) return false;
            }
            return true;
        },
        async download() {
            if (this.selectedSession.length === 0) {
                this.setErrorMessage(`session is not defined`)
                return;
            }
            if (this.urls.length === 0) {
                this.setErrorMessage(`url list empty`)
                return;
            }
            this.downloadSuccess = false;
            this.downloadFailure = false;
            this.setErrorMessage(``)
            if (this.fileIdStrategy === "pipe") {
                this.validateIds();
                if (this.hasErrors()) {
                    return;
                }
            }

            try {
                await this.prepareUrls().forEach((url) => {
                    enqueueDownload(this.selectedSession, url[1], url[0])
                })
                this.downloadSuccess = true;
            } catch (e) {
                this.downloadFailure = true;
            }
        },
        hasErrors() {
            return this.errorMessage.length > 0;
        },
        validateIds() {
            this.urls.split("\n").forEach((url, key) => {
                const parts = url.split(this.idUrlSeparator);
                if (parts.length !== 2) {
                    this.setErrorMessage(`line number ${key + 1} is not properly formatted, expected "id|url"`)
                }
            })
        },
        prepareUrls() {
            const urls = [];
            this.urls.split("\n").forEach((url, key) => {
                if (this.fileIdStrategy === "pipe") {
                    const parts = url.split(this.idUrlSeparator);
                    urls.push([parts[0], parts[1]])
                } else {
                    urls.push([this.fileIdStrategy, url])
                }
            })
            return urls;
        },
        setErrorMessage(s) {
            this.errorMessage = s
        }
    },
    async mounted() {
        this.setSessions(await getSessions())
    }
}