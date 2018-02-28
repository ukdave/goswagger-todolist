var myApp = new Vue({
    el: '#myApp',
    data: {
        items: [],
        newItem: {
            description: "",
            completed: false
        }
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
        },
        showAddItemModal: function() {
            this.newItem.description = "";
            this.newItem.completed = false;
            $('#addItemModal').modal('show');
        },
        submitForm: function () {
            var self = this;
            axios.post('/api/', self.newItem).then(function (response) {
                self.items.push(response.data);
                $('#addItemModal').modal('hide');
            });
        }
    }
});
