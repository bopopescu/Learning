Dataproc Quickstart
===================

.. _google-setup:

Configuring Google Cloud Platform (GCP) credentials
---------------------------------------------------

Configuring your GCP credentials allows mrjob to run your jobs on
Dataproc and use GCS.

* `Create a Google Cloud Platform account <http://cloud.google.com/>`_, see top-right
* `Learn about Google Cloud Platform "projects" <https://cloud.google.com/docs/overview/#projects>`_
* `Select or create a Cloud Platform Console project <https://console.cloud.google.com/project>`_
* `Enable billing for your project <https://console.cloud.google.com/billing>`_
* Go to the `API Manager <https://console.cloud.google.com/apis>`_ and search for / enable the following APIs...

  * Google Cloud Storage
  * Google Cloud Storage JSON API
  * Google Cloud Dataproc API

* Under Credentials, **Create Credentials** and select **Service account key**.  Then, select **New service account**, enter a Name and select **Key type** JSON.

* Install the `Google Cloud SDK <https://cloud.google.com/sdk/>`_

`Dataproc Documentation <https://cloud.google.com/dataproc/overview>`_

`How GCP Default credentials work <https://developers.google.com/identity/protocols/application-default-credentials#howtheywork>`_


.. _running-a-dataproc-job:

Running a Dataproc Job
----------------------

Running a job on Dataproc is just like running it locally or on your own Hadoop
cluster, with the following changes:

* The job and related files are uploaded to GCS before being run
* The job is run on Dataproc (of course)
* Output is written to GCS before mrjob streams it to stdout locally
* The Hadoop version is specified by the `Dataproc version <https://cloud.google.com/dataproc/dataproc-versions>`_

This the output of this command should be identical to the output shown in
:doc:`quickstart`, but it should take much longer::

    > python word_count.py -r dataproc README.txt
    "chars" 3654
    "lines" 123
    "words" 417

Sending Output to a Specific Place
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

If you'd rather have your output go to somewhere deterministic on GCS, which you
probably do, use :option:`--output-dir`::

    > python word_count.py -r dataproc README.rst \
    >   --output-dir=gs://my-bucket/wc_out/

It's also likely that since you know where your output is on GCS, you don't want
output streamed back to your local machine. For that, use
:option:`-no-output`::

    > python word_count.py -r dataproc README.rst \
    >   --output-dir=gs://my-bucket/wc_out/ \
    >   --no-output

There are many other ins and outs of effectively using mrjob with Dataproc.
This is a strictly no-outs body of documentation!

.. _picking-dataproc-cluster-config:

Choosing Type and Number of GCE Instances
-----------------------------------------

When you create a cluster on Dataproc, you'll have the option of specifying a number
and type of GCE instances, which are basically virtual machines. Each instance
type has different memory, CPU, I/O and network characteristics, and costs
a different amount of money. See
`Machine Types <https://cloud.google.com/compute/docs/machine-types>`_ and
`Pricing <https://cloud.google.com/compute/pricing>`_ for details.

Instances perform one of three roles:

* **Master**: There is always one master instance. It handles scheduling of tasks
  (i.e. mappers and reducers), but does not run them itself.
* **Worker**: You may have one or more worker instances. These run tasks and host
  HDFS.
* **Preemptible Worker**: You may have zero or more of these. These run tasks, but do *not*
  host HDFS. This is mostly useful because your cluster can lose task instances
  without killing your job (see `Preemptible VMs <https://cloud.google.com/dataproc/preemptible-vms>`_).

By default, :py:mod:`mrjob` runs a single ``n1-standard-1``, which is a cheap but not
very powerful instance type. This can be quite adequate for testing your code on a small subset of your
data, but otherwise give little advantage over running a job locally. To get more performance out of
your job, you can either add more instances, use more powerful instances, or both.

Here are some things to consider when tuning your instance settings:

* Google Cloud bills you a 10-minute minimum even if your cluster only lasts for a few
  minutes (this is an artifact of the Google Cloud billing structure), so for many
  jobs that you run repeatedly, it is a good strategy to pick instance settings
  that make your job consistently run in a little less than 10 minutes.
* Your job will take much longer and may fail if any task (usually a reducer)
  runs out of memory and starts using swap. (You can verify this by using
  :command:`vmstat`.) Restructuring your
  job is often the best solution, but if you can't, consider using a high-memory
  instance type.
* Larger instance types are usually a better deal if you have the workload
  to justify them. For example, a ``n1-highcpu-8`` costs about 6 times as much
  as an ``n1-standard-1``, but it has about 8 times as much processing power
  (and more memory).

The basic way to control type and number of instances is with the
:mrjob-opt:`instance_type` and :mrjob-opt:`num_core_instances` options, on the command line like
this::

    --instance-type n1-highcpu-8 --num-core-instances 4

or in :py:mod:`mrjob.conf`, like this::

    runners:
      dataproc:
        instance_type: n1-highcpu-8
        num_core_instances: 4

In most cases, your master instance type doesn't need to be larger
than ``n1-standard-1`` to schedule tasks.  *instance_type* only applies to
instances that actually run tasks. (In this example, there are 1 ``n1-standard-1``
master instance, and 4 ``n1-highcpu-8`` worker instances.) You *will* need a larger
master instance if you have a very large number of input files; in this case,
use the :mrjob-opt:`master_instance_type` option.

If you want to run preemptible instances, use the :mrjob-opt:`task_instance_type` and :mrjob-opt:`num_task_instances` options.


.. _dataproc-limitations:

Limitations
-----------

mrjob's Dataproc implementation is relatively new and does not yet have some
features supported by other runners, including:

* fetching counters
* finding probable cause of errors
* running Java JARs as steps
* :mrjob-opt:`libjars` support
