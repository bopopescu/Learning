Dataproc runner options
=======================

All options from :doc:`configs-all-runners` and :doc:`configs-hadoopy-runners`
are available to Dataproc runner.

Google credentials
------------------

See :ref:`google-setup` for specific instructions
about setting these options.


Choosing/creating a cluster to join
------------------------------------

.. mrjob-opt::
    :config: cluster_id
    :switch: --cluster-id
    :type: :ref:`string <data-type-string>`
    :set: dataproc
    :default: automatically create a cluster and use it

    The ID of a persistent Dataproc cluster to run jobs in.  It's fine for other
    jobs to be using the cluster; we give our job's steps a unique ID.


Cluster creation and configuration
-----------------------------------

.. mrjob-opt::
    :config: zone
    :switch: --zone
    :type: :ref:`string <data-type-string>`
    :set: dataproc
    :default: gcloud SDK default

    Availability zone to run the job in

.. mrjob-opt::
    :config: region
    :switch: --region
    :type: :ref:`string <data-type-string>`
    :set: dataproc
    :default: gcloud SDK default

    region to run Dataproc jobs on (e.g.  ``us-central-1``). Also used by mrjob
    to create temporary buckets if you don't set :mrjob-opt:`cloud_tmp_dir`
    explicitly.

.. mrjob-opt::
    :config: image_version
    :switch: --image-version
    :type: :ref:`string <data-type-string>`
    :set: dataproc
    :default: ``1.0``

    Cloud Image to run Dataproc jobs on.  See `the Dataproc docs on specifying the Dataproc version`_.  for details.

    .. _`the Dataproc docs on specifying the Dataproc version`:
        https://cloud.google.com/dataproc/dataproc-versions

Bootstrapping
-------------

These options apply at *bootstrap time*, before the Hadoop cluster has
started. Bootstrap time is a good time to install Debian packages or compile
and install another Python binary.

.. mrjob-opt::
    :config: bootstrap
    :switch: --bootstrap
    :type: :ref:`string list <data-type-string-list>`
    :set: dataproc
    :default: ``[]``

    A list of lines of shell script to run once on each node in your cluster,
    at bootstrap time.

    Passing expressions like ``path#name`` will cause
    *path* to be automatically uploaded to the task's working directory
    with the filename *name*, marked as executable, and interpolated into the
    script by their absolute path on the machine running the script. *path*
    may also be a URI, and ``~`` and environment variables within *path*
    will be resolved based on the local environment. *name* is optional.
    For details of parsing, see :py:func:`~mrjob.setup.parse_setup_cmd`.

    Unlike with :mrjob-opt:`setup`, archives are not supported (unpack them
    yourself).

    Remember to put ``sudo`` before commands requiring root privileges!


.. mrjob-opt::
   :config: bootstrap_python
   :switch: --bootstrap-python, --no-bootstrap-python
   :type: boolean
   :set: dataproc
   :default: ``True``

   Attempt to install a compatible version of Python at bootstrap time,
   including :command:`pip` and development libraries (so you can build
   Python packages written in C).

   This is useful even in Python 2, which is installed by default, but without
   :command:`pip` and development libraries.

Monitoring the cluster
-----------------------

.. mrjob-opt::
    :config: check_cluster_every
    :switch: --check-cluster-every
    :type: :ref:`string <data-type-string>`
    :set: dataproc
    :default: 10

    How often to check on the status of Dataproc jobs in seconds. If you set this
    too low, GCP will throttle you.

Number and type of instances
----------------------------

.. mrjob-opt::
    :config: instance_type
    :switch: --instance-type
    :type: :ref:`string <data-type-string>`
    :set: dataproc
    :default: ``'n1-standard-1'``

    What sort of GCE instance(s) to use on the nodes that actually run tasks
    (see https://cloud.google.com/compute/docs/machine-types).  When you run multiple
    instances (see :mrjob-opt:`instance_type`), the master node is just
    coordinating the other nodes, so usually the default instance type
    (``n1-standard-1``) is fine, and using larger instances is wasteful.

.. mrjob-opt::
    :config: master_instance_type
    :switch: --master-instance-type
    :type: :ref:`string <data-type-string>`
    :set: dataproc
    :default: ``'n1-standard-1'``

    like :mrjob-opt:`instance_type`, but only for the master Hadoop node.
    This node hosts the task tracker and HDFS, and runs tasks if there are no
    other nodes. Usually you just want to use :mrjob-opt:`instance_type`.

.. mrjob-opt::
    :config: core_instance_type
    :switch: --core-instance-type
    :type: :ref:`string <data-type-string>`
    :set: dataproc
    :default: value of :mrjob-opt:`instance_type`

    like :mrjob-opt:`instance_type`, but only for worker Hadoop nodes; these nodes run tasks and host HDFS. Usually you
    just want to use :mrjob-opt:`instance_type`.


.. mrjob-opt::
    :config: task_instance_type
    :switch: --task-instance-type
    :type: :ref:`string <data-type-string>`
    :set: dataproc
    :default: value of :mrjob-opt:`instance_type`

    like :mrjob-opt:`instance_type`, but only for the task Hadoop nodes;
    these nodes run tasks but do not host HDFS. Usually you just want to use
    :mrjob-opt:`instance_type`.


.. mrjob-opt::
    :config: num_core_instances
    :switch: --num-core-instances
    :type: :ref:`string <data-type-string>`
    :set: dataproc
    :default: 2

    Number of worker instances to start up. These run your job and
    host HDFS.

.. mrjob-opt::
    :config: num_task_instances
    :switch: --num-task-instances
    :type: :ref:`string <data-type-string>`
    :set: dataproc
    :default: 0

    Number of task instances to start up.  These run your job but do not host
    HDFS. If you use this, you must set :mrjob-opt:`num_core_instances`; Dataproc does not allow you to
    run task instances without core instances (because there's nowhere to host
    HDFS).

FS paths and options
--------------------
MRJob uses google-api-python-client to manipulate/access FS.

.. mrjob-opt::
    :config: cloud_tmp_dir
    :switch: --cloud-tmp-dir
    :type: :ref:`string <data-type-string>`
    :set: dataproc
    :default: (automatic)

    GCS directory (URI ending in ``/``) to use as temp space, e.g.
    ``gs://yourbucket/tmp/``.

    By default, mrjob looks for a bucket belong to you whose name starts with
    ``mrjob-`` and which matches :mrjob-opt:`region`. If it can't find
    one, it creates one with a random name. This option is then set to `tmp/`
    in this bucket (e.g. ``gs://mrjob-01234567890abcdef/tmp/``).

.. mrjob-opt::
    :config: cloud_fs_sync_secs
    :switch: --cloud-fs-sync-secs
    :type: :ref:`string <data-type-string>`
    :set: dataproc
    :default: 5.0

    How long to wait for GCS to reach eventual consistency. This is typically
    less than a second, but the default is 5.0 to be safe.
