import AuthRoute from './auth';

export default AuthRoute.extend({
	model: function(args) {
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
  	},
    saveStep:function() {
      alert('saveStep');
      console.log('saving step');
    },
    rollbackStep:function() {
      alert('rollbackStep');
      console.log('rollback step');
    }
});
