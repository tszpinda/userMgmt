import Ember from 'ember';

export default Ember.Controller.extend({

    isEditing: false,
    isReadOnly : function() {
        return !this.get('isEditing');
    }.property('isEditing'),

    isShowingConfirmDelete: false,



    actions: {
        confirmDelete:function() {
            this.toggleProperty('isShowingConfirmDelete');
        },
        deleteConfirmed : function() {
            this.get('model').destroyRecord();
            this.transitionToRoute('tutorials');
            this.toggleProperty('isShowingConfirmDelete');
        },
        toggleEditing: function(){
            this.toggleProperty('isEditing');
        },
        updateTutorial: function() {
            var _this = this;
            this.model.get('errors').clear();
            this.model.save().then(function(model) {
                //_this.transitionTo('tutorial.edit', model);
                console.log("tutorial updated", model.get('name'));
                _this.toggleProperty('isEditing');
            }, function() {
                _this.notifier.error("Updating a tutorial failed");
            });
        },
    }

});
