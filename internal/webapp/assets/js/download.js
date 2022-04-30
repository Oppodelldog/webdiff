const componentDownload = {
    template: `
      <div class="row">
      <h6>Download</h6>
        </div>`,
    data() {
        return {}
    },
    methods: {},
    async mounted() {
        this.setSessions(await getSessions())
    }
}