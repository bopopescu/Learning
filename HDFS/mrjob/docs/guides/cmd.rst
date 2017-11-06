.. _mrjob-cmd:

The ``mrjob`` command
=====================

The :command:`mrjob` command has two purposes:

1. To provide easy access to EMR tools
2. To eventually let you run Hadoop Streaming jobs written in languages other
   than Python

The :command:`mrjob` command comes with Python-version-specific aliases (e.g.
:command:`mrjob-3`, :command:`mrjob-3.4`), in case you choose to install
``mrjob`` for multiple versions of Python.

EMR tools
---------

.. _audit-emr-usage:

audit-emr-usage
^^^^^^^^^^^^^^^

   .. automodule:: mrjob.tools.emr.audit_usage

boss
^^^^

    .. automodule:: mrjob.tools.emr.mrboss

create-cluster
^^^^^^^^^^^^^^

   .. automodule:: mrjob.tools.emr.create_cluster

.. _report-long-jobs:

report-long-jobs
^^^^^^^^^^^^^^^^

    .. automodule:: mrjob.tools.emr.report_long_jobs

s3-tmpwatch
^^^^^^^^^^^

    .. automodule:: mrjob.tools.emr.s3_tmpwatch

terminate-cluster
^^^^^^^^^^^^^^^^^

    .. automodule:: mrjob.tools.emr.terminate_cluster

.. _terminate-idle-clusters:

terminate-idle-clusters
^^^^^^^^^^^^^^^^^^^^^^^

    .. automodule:: mrjob.tools.emr.terminate_idle_clusters


Running jobs
------------

    ``mrjob run (path to script or executable) [options]``

    Run a job. Takes same options as invoking a Python job. See
    :doc:`configs-all-runners`, :doc:`configs-hadoopy-runners`, :doc:`dataproc-opts`, and
    :doc:`emr-opts`. While you can use this command to invoke your jobs, you
    can just as easily call ``python my_job.py [options]``.
