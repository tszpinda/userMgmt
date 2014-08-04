/*
The Session model stores info about the current session, including
the current user. The current user may not be present if it is a
logged out session (i.e. there is always a current session,
regardless of whether a user has logged in or not)
 */
import DS from "ember-data";

var Session = DS.Model.extend({
  email:      DS.attr('string'),
  password:   DS.attr('string'),
  authToken:  DS.attr('string'),
  user:       DS.belongsTo('user'),
  csrfToken:  DS.attr('string')
});

export default Session;
