import Ember from 'ember';

export default Ember.ObjectController.extend({

	actions: { 
        saveStep:function(step) {
          console.log('saving step', step.get('text'));
          step.save();
        },
        rollbackStep:function(step) {
          console.log('rollback step', step.get('text'));
          step.rollback();
        },
        selectElement:function(step) {
          console.log('selectElement', step.get('text'));
            debugger
          chrome.runtime.sendMessage({selectElement: true,
                                      tutorialId: this.get('model').get('id'),
                                      stepId: step.get('id'),
                                      stepText: step.get('text')}, function(response) {
            console.log("Response in tutorial.js:", response);
          });
          window.close();
        }
    }

})