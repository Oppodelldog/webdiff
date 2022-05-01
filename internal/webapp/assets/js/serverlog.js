const componentServerLog = {
    template: `
   <div aria-live="polite" aria-atomic="true" class="">

        <div class="toast-container position-absolute top-0 end-0 p-3">

            <div v-for="(message,idx) in messages" class="toast show" role="alert" aria-live="assertive" aria-atomic="true">
                <div class="toast-header">
                    <strong class="me-auto">{{message.title}}</strong>
                    <small class="text-muted">{{message.severity}}</small>
                    <button type="button" class="btn-close" aria-label="Close" @click="closeToast(idx)"></button>
                </div>
                <div class="toast-body">
                    {{message.message}}
                </div>
            </div>
            
        </div>
    </div>
`,
    data() {
        return {
            messages: []
        }
    },
    methods: {
        newServerLogMessage(message) {
            this.messages.push({title: "Server Log", severity: message.severity, message: message.message,})
        },
        closeToast(idx) {
            this.messages.splice(idx, 1);
        }
    },
    async mounted() {
        ServerLogObserver.addListener(this.newServerLogMessage)
    },
    unmounted() {
        ServerLogObserver.removeListener(this.newServerLogMessage)
    }
}