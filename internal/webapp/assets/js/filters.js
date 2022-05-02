const componentFilters = {
    template: `
<div class="accordion" id="accordionExample">
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
                    <span class="input-group-text">New</span>
                    <input type="text" class="form-control" placeholder="Filter name" aria-label="Filter"
                           v-model="selectedFilterName" @input="filterSelected()">
                    <span class="input-group-text">Existing</span>
                    <select class="form-select" aria-label="Default filter" v-model="selectedFilterName"
                            @change="filterSelected()">
                        <option disabled value="">Select a filter</option>
                        <option v-for="filter in filters" :value="filter.name">{{filter.name}}</option>
                    </select>
                </div>

                <div class="input-group mt-3">
                    <input type="text" class="form-control" placeholder="Filter" aria-label="Filter"
                           v-model="selectedFilterFilter">
                </div>
                <div class="mt-3">
                    <button type="button" class="btn btn-primary" @click="saveFilter">Save</button>
                    <button type="button" class="btn btn-danger ms-3" v-if="isExistingFilterSelected()"
                            @click="deleteFilter">Delete
                    </button>
                </div>
            </div>
        </div>
    </div>
</div>
`,
    data() {
        return {
            filters: [],
            selectedFilterName: this.modelValue,
            selectedFilterFilter: "",
        }
    },
    props: {
        modelValue: {
            type: String,
            default: ''
        }
    },
    methods: {
        async saveFilter() {
            await upsertFilter(this.selectedFilterName, this.selectedFilterFilter)
            await this.loadFilters();
            this.notify()
            this.$emit('saved');
        },
        async deleteFilter() {
            await deleteFilter(this.selectedFilterName)
            await this.loadFilters();
            this.selectedFilterFilter = ""
            this.selectedFilterName = "";
            this.notify()
            this.$emit('deleted');
        },
        notify: function () {
            this.$emit('update:modelValue', this.selectedFilterName);
        },
        filterSelected() {
            for (let i = 0; i < this.filters.length; i++) {
                if (this.filters[i].name === this.selectedFilterName) {
                    this.selectedFilterFilter = this.filters[i].filter;
                    this.notify();
                    return;
                }
            }
            this.selectedFilterFilter = "";
        },
        isExistingFilterSelected() {
            for (let i = 0; i < this.filters.length; i++) {
                if (this.filters[i].name === this.selectedFilterName) return true;
            }
            return false;
        },
        async loadFilters() {
            this.filters = (await getFilters()).filters;
        },
    },
    async mounted() {
        await this.loadFilters();
    }
}