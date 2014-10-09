import AuthRoute from './../auth';

export default AuthRoute.extend({
	model: function() {
    console.log("route:tutorial:new");
  	return this.store.createRecord('tutorial');
 	},

  actions: {
   	createTutorial: function() {
    	var _this = this;

     	this.currentModel.get('errors').clear();
     	this.currentModel.save().then(function(model) {
          console.log("saved", model.get('name'));
        	//_this.transitionTo('tutorial.edit', model);
         _this.transitionTo('tutorial.edit', _this.currentModel);
     	}, function() {
       	_this.notifier.error("Creating a new tutorial failed");
     	});
    }
  }
});