import AuthRoute from './auth';

export default AuthRoute.extend({
	model: function(args) {
		console.log("route:tutorial", args);
		return this.store.find('tutorial', args.tutorial_id);
	},

	actions: {

        saveStep:function(step) {
            console.log('route:tutorial: saving step', step.get('text'));
            step.save();
        },
        rollbackStep:function(step) {
            console.log('route:tutorial: rollback step', step.get('text'));
            step.rollback();
        },
        deleteStep:function(step) {
            console.log('route:tutorial: delete step', step.get('text'));
            step.destroyRecord();
        },
        addStep: function() {
            var step = this.store.createRecord('step');
            step.set('tutorial', this.currentModel);
            this.currentModel.get('steps').add(step);
        }
  	}
});
