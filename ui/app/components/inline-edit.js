import Ember from 'ember';

export default Ember.Component.extend({


    hasValue : function() {
        var v = this.get('value');
        if (v === 0){
            return true;
        }
      return this.get('value');
    }.property('value'),

	actions: {
		save: function(){
			console.log('saving');
			this.toggleProperty('isEditing');
			this.sendAction('onSave', this.get('item'));
		},
		toggleEditing: function(){
			console.log('toggleEditing');
			if(this.get('isEditing')) {
                this.sendAction('onCancel', this.get('item'));
            }
			this.toggleProperty('isEditing');
		}
	}
});


