import AuthRoute from './auth';

export default AuthRoute.extend({
	model: function(args) {
		//alert('route:tutorial');
		console.log("route:step", args);
		return this.store.find('step', args.step_id);
	},

	actions: {
   		updateStep: function() {
    		var _this = this;
     		this.currentModel.get('errors').clear();
     		this.currentModel.save().then(function(model) {
        		//_this.transitionTo('tutorial.edit', model);
         		console.log("step updated", model.get('selector'));
         		_this.transitionTo('steps');
     		}, function() {
       			_this.notifier.error("Updating a step failed");
     		});
    	}
  	}
});
