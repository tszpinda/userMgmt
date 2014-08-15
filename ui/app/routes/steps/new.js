import AuthRoute from './../auth';

export default AuthRoute.extend({
	model: function() {
    	console.log("route:steps:new");
  		var step = this.store.createRecord('step');
  		step.set('tutorial', this.modelFor('tutorial'));
  		return step;
 	},

  actions: {
   	createStep: function() {
    	var _this = this;
    	var step  = this.currentModel;
     	step.get('errors').clear();
     	console.log('Adding step with selector', step.get('selector'));
     	//var t = step.get('tutorial');
     	//t.get('steps').add(step)
     	step.save().then(function(model) {
          console.log("saved", model.get('selector'));
        	//_this.transitionTo('tutorial.edit', model);
         _this.transitionTo('steps');
     	}, function() {
       	_this.notifier.error("Creating a new step failed");
     	});
    }
  }
});
