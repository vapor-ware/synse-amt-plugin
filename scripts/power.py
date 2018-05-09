# Python commands for amt power status and management
import sys
from amt.client import Client

def main():
    """Communicate with a given amt device
    """
    ip = sys.argv[1]
    password = sys.argv[2]
    command_name = sys.argv[3]

    amt_client = Client(ip, password)

    commands = {
        'status': amt_client.power_status,
        'on': amt_client.power_on,
        'off': amt_client.power_off,
        'cycle': amt_client.power_cycle
        }

    # Not sure exactly where these codes are defined yet, still looking
    status_codes = {
        '2': 'on',
        '8': 'off'
    }

    if command_name in commands:
        output = commands.get(command_name)()
        if command_name == 'status':
            output = status_codes.get(output)
            if not output:
                raise ValueError('Unexpected status "{}" returned for amt power status', output)
            print(output)
    else:
        raise ValueError('No command {} found for amt power device'.format(command_name))

if __name__ == '__main__':
    main()
