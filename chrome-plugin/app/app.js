import Ember from 'ember';
import Resolver from 'ember/resolver';
import loadInitializers from 'ember/load-initializers';

Ember.MODEL_FACTORY_INJECTIONS = true;

var App = Ember.Application.extend({
  modulePrefix: 'ember-chrome-demo', // TODO: loaded via config
  Resolver: Resolver,
  /*
  LOG_TRANSITIONS: true,
	LOG_BINDINGS: true,
	LOG_VIEW_LOOKUPS: true,
	LOG_STACKTRACE_ON_DEPRECATION: true,
	LOG_VERSION: true,
	LOG_TRANSITIONS_INTERNAL: true,
	debugMode: true*/
});

loadInitializers(App, 'ember-chrome-demo');

export default App;
