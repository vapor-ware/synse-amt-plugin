# -*- coding: utf-8 -*-

"""Python commands for amt boot target selection"""

import sys
from amt.client import Client


def main():
    """Communicate with amt device
    """
    host = sys.argv[1]
    password = sys.argv[2]
    boot_target = sys.argv[3]

    amt_client = Client(host, password)
    amt_client.set_next_boot(boot_target)
    

if __name__ == '__main__':
    main()
