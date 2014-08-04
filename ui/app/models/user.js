import DS from "ember-data";

var User = DS.Model.extend({
  name:                 DS.attr('string'),
  password:             DS.attr('string'),
  passwordConfirmation: DS.attr('string'),
  email:                DS.attr('string')
});

export default User;