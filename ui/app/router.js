import Ember from 'ember';

var Router = Ember.Router.extend({
  location: UserMgmtENV.locationType
});

Router.map(function() {
  this.route('signup');
  this.route('login');  	
  this.route('dashboard');  
  //this.resource('tutorial', { path: 'tutorials' });
  //this.route('tutorial/edit')
  //this.resource('tutorial/step', { path: 'tutorial/steps/:tutorial/step_id' });


//  this.route('tutorial/new');
  this.resource('tutorials', function(){
  	this.route('new');

    this.resource('tutorial', { path: '/:tutorial_id' }, function() {
      this.route('edit');
    });

  });
});

export default Router;
