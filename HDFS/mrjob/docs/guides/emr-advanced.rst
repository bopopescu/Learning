Advanced EMR usage
==================

.. _spot-instances:

Spot Instances
--------------

You can potentially save money purchasing EC2 instances for your EMR
clusters from AWS's spot market. The catch is that if someone bids more for
instances that you're using, they can be taken away from your cluster. If this
happens, you aren't charged, but your job may fail.

You can specify spot market bid prices using the *core_instance_bid_price*,
*master_instance_bid_price*, and *task_instance_bid_price* options to
specify a price in US dollars. For example, on the command line::

    --ec2-task-instance-bid-price 0.42

or in :py:mod:`mrjob.conf`::

    runners:
      emr:
        task_instance_bid_price: '0.42'

(Note the quotes; bid prices are strings, not floats!)

Amazon has a pretty thorough explanation of why and when you'd want to use spot
instances `here
<http://docs.amazonwebservices.com/ElasticMapReduce/latest/DeveloperGuide/UsingEMR_SpotInstances.html?r=9215>`_.
The brief summary is that either you don't care if your job fails, in which
case you want to purchase all your instances on the spot market, or you'd need
your job to finish but you'd like to save time and money if you can, in which
case you want to run task instances on the spot market and purchase master and
core instances the regular way.

:ref:`cluster-pooling` interacts with bid prices more or less how you'd
expect; a job will join a pool with spot instances only if it requested spot
instances at the same price or lower.

Custom Python packages
----------------------

See :ref:`using-pip` and :ref:`installing-packages`.

.. _bootstrap-time-configuration:

Bootstrap-time configuration
----------------------------

Some Hadoop options, such as the maximum number of running map tasks per node,
must be set at bootstrap time and will not work with `--jobconf`. You must use
Amazon's `configure-hadoop` script for this. For example, this limits the
number of mappers and reducers to one per node::

    --bootstrap-action="s3://elasticmapreduce/bootstrap-actions/configure-hadoop \
    -m mapred.tasktracker.map.tasks.maximum=1 \
    -m mapred.tasktracker.reduce.tasks.maximum=1"

.. note::

   This doesn't work on AMI 4.0.0 and later.

.. _reusing-clusters:

Manually Reusing Clusters
-------------------------

In some cases, it may be useful to have more fine-grained control than
:ref:`cluster-pooling` provides; for example, to run several related jobs
on the same cluster.

.. warning::

   If you do this on mrjob versions prior to 0.6.0, make sure to set
   :mrjob-opt:`max_hours_idle`, or your manually created clusters will
   run forever, costing you money.

:py:mod:`mrjob` includes a utility to create persistent clusters without
running a job. For example, this command will create a cluster with 12 EC2
instances (1 master and 11 core), taking all other options from
:py:mod:`mrjob.conf`::

    $ mrjob create-cluster --num-core-instances=11 --max-hours-idle 1
    ...
    j-CLUSTERID

You can then add jobs to the cluster with the :option:`--emr-cluster-id`
switch or the `emr_cluster_id` variable in `mrjob.conf` (see
:py:meth:`EMRJobRunner.__init__`)::

    $ python mr_my_job.py -r emr --emr-cluster-id=j-CLUSTERID input_file.txt > out
    ...
    Adding our job to existing cluster j-CLUSTERID
    ...

Debugging will be difficult unless you complete SSH setup (see
:ref:`ssh-tunneling`) since the logs will not be copied from the master node to
S3 before either five minutes pass or the cluster terminates.
