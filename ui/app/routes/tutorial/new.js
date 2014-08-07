import Ember from 'ember';

export default Ember.Route.extend({
	model: function() {
    	return this.store.createRecord('tutorial');
  	},

  	actions: {
    	createTutorial: function() {
      	var _this = this;

      	this.currentModel.get('errors').clear();

      	this.currentModel.save().then(function(model) {
        	_this.transitionTo('tutorial', model);
      	}, function() {
        	_this.notifier.error("Creating a new tutorial failed");
      	});
    }
  }
});
