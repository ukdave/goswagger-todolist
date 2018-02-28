var myApp = new Vue({
    el: '#myApp',
    data: {
        items: []
    },
    computed: {
        orderedItems: function () {
            return _.sortBy(this.items, 'id')
        }
    },
    created: function () {
        this.refreshItems();
    },
    methods: {
        refreshItems: function () {
            var self = this;
            axios.get('/api/').then(function (response) {
                self.items = response.data
            });
        }
    }
});
