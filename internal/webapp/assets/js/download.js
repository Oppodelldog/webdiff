const componentDownload = {
    template: `
 <div class="row">
    <div class="input-group mb-3">
        <div class="mb-3">
            <label for="floatingTextarea2">Session</label>
            <input type="text" class="form-control" placeholder="New Session" aria-label="New Session"
                   v-model="selectedSession">
        </div>
        <div class="mb-3">
            <label for="floatingTextarea2">existing sessions</label>
            <select v-model="selectedSession" class="form-select"
                    aria-label="Choose a session to load file for selection on the left side">
                <option disabled value="">Choose a session</option>
                <option v-for="session in sessions" :value="session">{{session}}</option>
            </select>
        </div>
    </div>
    <div class="mb-3">
        <label for="floatingTextarea2">Urls</label>
        <textarea v-model="urls" class="form-control" placeholder="one url per line" id="urlList"
                  style="height: 200px"></textarea>
    </div>
    <div class="row">

        <div class="col">
            <button type="button" :disabled="isFillUrlsButtonDisabled()" @click="fillUrls()" class="btn btn-outline-primary">Fill Urls from
                existing session
            </button>
        </div>
        <div class="col">
        </div>
        <div class="col-auto">
            <button type="button" class="btn btn-primary" @click="download()">Download</button>
        </div>
    </div>
</div>
        `,
    data() {
        return {
            sessions: [],
            urls: "",
            selectedSession: "",
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
            if (this.selectedSession.length === 0) return;
            if (this.urls.length === 0) return;

            await this.urls.split("\n").forEach((url) => {
                enqueueDownload(this.selectedSession, url)
            })
        },
    },
    async mounted() {
        this.setSessions(await getSessions())
    }
}