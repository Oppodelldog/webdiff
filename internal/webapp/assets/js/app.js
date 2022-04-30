let app = Vue.createApp({
    data() {
        return {
            sessions: [],
            filesA: [],
            filesB: [],
            selectedSessionA: "",
            selectedSessionB: "",
            selectedFilesA: [],
            selectedFilesB: [],
            diffContents: ["select files to diff"]
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
        },
        async loadFilesB() {
            if (this.selectedSessionB.length === 0) return;

            const data = await getFiles(this.selectedSessionB)
            this.filesB = data.files;
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

                const diffResult = await getDiff(sessionA, fileA, sessionB, fileB)
                this.diffContents.push(diffResult.diff)
            }
        }
    },
    async mounted() {
        this.setSessions(await getSessions())
    }
});

app.mount('#app')
