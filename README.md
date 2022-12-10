2022.12.10 brainstorming:

- a CLI application that takes in the kubeconfig, and has a nice interactive CLI menu where you can choose if you want to do a mutating webhook or a validating webhook. Then, you can select which type of object you want to create the admission controller rule for (POD,DaemonSet,Deployment, etc.). Then, you select what you want (e.g resource limit, label, selector, tag, etc.), type it in, press enter, and bumm magic happens.

so something like:

[MUTATING or VALIDATING] -> [RESOURCE (pod,deployment, etc)] -> [OPERATION (create,delete,etc.)] -> (the PATH of the admission controller will come from this, usually if there are already rules for pods, it will be the same) -> [THE RULE (e.g resource limit must be under 200Mi)] -> SAVE.

- You should be able to see the already existing rules for the cluster, so for that we'll need a database. Multiple people can create admission controllers for the same cluster. problem: with a CLI app you cannot really see who's administering the controller (you can get the linux username, but if someone is root...). This brings up the question whether this should be a web application, or stay as a CLI one. Also if it's a CLI one, how would I effectively share the DB between users....

I shouldn't let perfect be the enemy of good, so it'll start out as a CLI app, and we'll see. We should be able to DELETE admission controller rules as well.

For an admission controller to work, we need:

- ValidatingWebhookConfiguration + MutatingWebhookConfiguration objects in K8s.
