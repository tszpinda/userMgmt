import Ember from 'ember';

var Router = Ember.Router.extend({
  location: UserMgmtENV.locationType
});

Router.map(function() {
  this.route('signup');
  this.route('login');  	
  this.route('dashboard');  
  this.resource('tutorial', { path: 'tutorials' });
  //this.route('tutorial/edit')
  //this.resource('tutorial/step', { path: 'tutorial/steps/:tutorial/step_id' });


  this.resource('tutorial/edit', { path: 'tutorials/:tutorial_id/edit' });
  this.route('tutorial/new');
});

export default Router;
