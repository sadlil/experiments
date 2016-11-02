#!/usr/bin/env node

"use strict";

//package declarations
var commandLineArgs = require('command-line-args');

//command line options
var cli = commandLineArgs([{
  name: 'sid',
  type: String
}, {
  name: 'auth',
  type: String
}]);

var options = cli.parse();

if (options.sid === undefined || options.auth === undefined) {
  console.log("Usage: twilio-media-delete --sid=[SID] --auth=[auth token]");
  process.exit(1);
}

//build our client
var client = require('twilio')(options.sid, options.auth);

client.messages.list({
  pageSize: 1000
}, function(err, data) {
  if (err) {
    console.log(err);
    console.log("This usually means your credentials are bad.");
    process.exit(1);
  } else {
    console.log(data.messages.length + " messages to delete (this may need to be more than once)");
    data.messages.forEach(function(message) {
      client.messages(message.sid).media.list(function(err, data) {
        if (data && data.media_list.length > 0) {
          data.media_list.forEach(function(mediaEntry, index) {
            console.log("Deleting Media SID " + mediaEntry.sid);
            client.messages(mediaEntry.parent_sid).media(mediaEntry.sid).delete(function(err, data) {
              if (err) {
                console.log("Media deletion error: " + err.status + " -- " + err.message);
              } else {
                console.log("Deleted Media SID " + mediaEntry.sid + " successfully.");

                if(!data){
                  console.log("Message SID " + message.sid + " media purged; deleting..." );
                  client.messages(message.sid).delete(function (err, data) {
                      if (err) {
                          console.log("Message deletion error: " + err.status + " -- " + err.message);
                      } else {
                          console.log("Deleted Message SID " + message.sid + " successfully.");
                      }
                  });
                }
              }
            });
          });
        } else {
          client.messages(message.sid).delete(function (err, data) {
              if (err) {
                  console.log("Message deletion error: " + err.status + " -- " + err.message);
              } else {
                  console.log("Deleted Message SID " + message.sid + " successfully.");
              }
          });
        }
      });
    });
  }
});