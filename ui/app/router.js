import Ember from 'ember';

var Router = Ember.Router.extend({
  location: UserMgmtENV.locationType
});

Router.map(function() {
  this.route('signup');
  this.route('login');  	
  this.route('dashboard');    
  this.resource('tutorials', function(){
  	this.route('new');

    this.resource('tutorial', { path: '/:tutorial_id' }, function() {
      this.route('edit');
      this.resource('steps', function(){
        this.route('new');        
        this.resource('step', { path: '/:step_id' }, function() {
          this.route('edit');
        });
      });
    });

  });
});

export default Router;
