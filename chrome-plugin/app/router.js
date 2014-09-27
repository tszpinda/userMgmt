import Ember from 'ember';

var Router = Ember.Router.extend({
//  location: 'none'
    location: 'hash'
});

Router.map(function() {
  this.route('signup');
  this.route('login');  	
  this.route('info');
  this.resource('tutorials', function(){
  	this.route('new');
      this.resource('tutorial', { path: '/:tutorial_id' }, function() {
        this.route('edit');
      });

  });
  //tutorial.edit - TutorialEditRoute
  //                TutorialEditController

  this.resource('tutorial', { path: '/tutorial/:tutorial_id' }, function(){
  	this.route('edit');
    this.route('view');
    /*this.resource('view',  function(){
        this.route('steps', { path: '/:tutorial_id' });
    });*/
    
  });
  this.route('bookmarks');
  this.route('tabs');
});

export default Router;
