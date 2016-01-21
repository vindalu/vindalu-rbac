vindalu-rbac
============

Role based access control for vindalu.


Policy
------
A `Policy` is a permission for a given path.

##### Name
The name of the policy.

##### Path
The requests URL path.  This is a regular expression representing a endpoint path.

##### Action
The action to be performed upon a successful match.  This can be the following 2 values:

	allow
	deny

##### Op
The operation being performed.  This can be any of the following:

	create
	read
	update
	delete
	all


Role
----
A `Role` contains a set of policy definitions.

##### Name
The name of the role.

##### Policies
An ordered list of policies for a given role.  The first matching `allow` (if found) is returned.


RoleMapping
-----------
A `RoleMapping` contains a user or group mapped to a set of roles.

##### Name
The name of a user or group.

##### Roles
An ordered list of roles assigned to a user or group.