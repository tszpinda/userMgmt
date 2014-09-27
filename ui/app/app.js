import Ember from 'ember';
import Resolver from 'ember/resolver';
import loadInitializers from 'ember/load-initializers';
import config from './config/environment';

Ember.MODEL_FACTORY_INJECTIONS = true;

var App = Ember.Application.extend({
  modulePrefix: config.modulePrefix,
  podModulePrefix: config.podModulePrefix,
  Resolver: Resolver,
  LOG_TRANSITIONS: true,
	LOG_BINDINGS: true,
	LOG_VIEW_LOOKUPS: true,
	LOG_STACKTRACE_ON_DEPRECATION: true,
	LOG_VERSION: true,
	//LOG_TRANSITIONS_INTERNAL: true,
	debugMode: true
});

loadInitializers(App, config.modulePrefix);

export default App;
