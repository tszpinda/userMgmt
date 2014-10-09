import AuthRoute from './auth';

export default AuthRoute.extend({
	model: function(args) {
		return this.store.find('tutorial', args.tutorial_id);
	},
    setupController: function(controller, model) {
        console.log("route:tutorial:setupController", model, controller);
        controller.set('model', model);
    },
	actions: {
        addStep:function() {
            console.log("adding step");
            var step = this.store.createRecord('step');
            step.set('tutorial', this.currentModel);
            //this.currentModel.get('steps').add(step);
            step.save();
        },
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
