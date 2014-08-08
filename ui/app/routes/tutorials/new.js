import AuthRoute from './../auth';

export default AuthRoute.extend({
	model: function() {
    	return this.store.createRecord('tutorial');
  	},

  	actions: {
    	createTutorial: function() {
      	var _this = this;

      	this.currentModel.get('errors').clear();

      	this.currentModel.save().then(function(model) {
        	//_this.transitionTo('tutorial.edit', model);
          _this.transitionTo('tutorials');
      	}, function() {
        	_this.notifier.error("Creating a new tutorial failed");
      	});
    }
  }
});