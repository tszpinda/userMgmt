import Ember from 'ember';

export default Ember.Component.extend({
	
	actions: {
		save: function(){
			console.log('saving');
			this.toggleProperty('isEditing');
			this.sendAction('onSave', this.get('item'));
		},
		toggleEditing: function(){
			console.log('toggleEditing');
			if(this.get('isEditing'))
				this.sendAction('onCancel', this.get('item'));
			this.toggleProperty('isEditing');
		}
	}
});


