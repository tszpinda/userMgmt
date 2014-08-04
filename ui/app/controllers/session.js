import Ember from 'ember';

/*
The SessionController wraps the session object, and can store a
transition so that you can retry it if you login - so if you load a
route that needs authentication but aren't logged in you can go back
 to that route after login
 */
export default Ember.ObjectController.extend({
	//needs:                  ['organization', 'organizations'],
  //currentOrganization:    Em.computed.alias('controllers.organization'),
  //organizations:          Em.computed.alias('controllers.organizations'),
  savedTransition:        null,
  isShowingOrganizations: false,

  actions: {
    toggleOrganizations: function() {
      this.toggleProperty('isShowingOrganizations');
    }
  }
});
