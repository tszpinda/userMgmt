import AuthRoute from './auth';

export default AuthRoute.extend({
	model: function(args) {
		//alert('route:tutorial');
		console.log("route:tutorial", args);
		return this.store.find('tutorial', args.tutorial_id);
	},

	actions: {
   		updateTutorial: function() {
    		var _this = this;
     		this.currentModel.get('errors').clear();
     		this.currentModel.save().then(function(model) {
        		//_this.transitionTo('tutorial.edit', model);
         		console.log("tutorial updated", model.get('name'));
         		_this.transitionTo('tutorials');
     		}, function() {
       			_this.notifier.error("Updating a tutorial failed");
     		});
    	}
  	}
});
